pragma solidity ^0.8.0;

contract CertificateManagement {
    
    struct Certificate {
        uint id; // 证书ID
        string studentName; // 学生姓名
        uint grade; // 成绩
        string course; // 课程
        address studentAddress; // 学生地址
        uint applicationTime; // 申请时间
        uint approvalTime; // 审核时间
        string status; // 证书状态
        address approver; // 审核人
    }
    
    mapping(uint => Certificate) public certificates; // 证书ID映射到证书详情
    mapping(address => uint[]) public studentCertificates; // 学生地址映射到其证书列表
    uint[] public pendingCertificates; // 待审核证书列表
    
    uint public certificateCount; // 证书总数
    
    // 事件定义
    event CertificateApplied(uint indexed id, address indexed studentAddress);
    event CertificateApproved(uint indexed id, address indexed approver);
    event CertificateRejected(uint indexed id, address indexed approver);
    
    // 申请证书
    function applyForCertificate(address _address, string memory _studentName, string memory _course) external {
        certificateCount++; // 递增证书总数
        certificates[certificateCount] = Certificate(certificateCount, _studentName, 0, _course, _address, block.timestamp, 0, "Pending", address(0)); // 创建新的证书
        studentCertificates[_address].push(certificateCount); // 更新学生的证书列表
        pendingCertificates.push(certificateCount); // 将证书加入待审核列表
        emit CertificateApplied(certificateCount, _address); // 触发事件
    }

    // 修改待审核的证书
    function modifyPendingCertificate(address _address, uint _id, string memory _studentName, string memory _course) external {
        require(keccak256(abi.encodePacked(certificates[_id].status)) == keccak256(abi.encodePacked("Pending")), "Certificate is not pending"); // 断言证书状态为待审核
        require(certificates[_id].studentAddress == _address, "Student address does not match"); // 断言学生地址匹配
        certificates[_id].studentName = _studentName; // 更新学生姓名
        certificates[_id].course = _course; // 更新课程
        emit CertificateApplied(_id, certificates[_id].studentAddress); // 触发事件
    }
    
    // 审核通过证书
    function approveCertificate(address _address, uint _id, uint _grade) external {
        require(keccak256(abi.encodePacked(certificates[_id].status)) == keccak256(abi.encodePacked("Pending")), "Certificate is not pending"); // 断言证书状态为待审核
        require(certificates[_id].approver == address(0), "Certificate already approved or rejected"); // 断言证书未被审核过
        certificates[_id].grade = _grade; // 更新证书成绩
        certificates[_id].status = "Approved"; // 更新证书状态为审核通过
        certificates[_id].approvalTime = block.timestamp; // 更新审核时间
        certificates[_id].approver = _address; // 更新审核人
        removePendingCertificate(_id); // 从待审核列表中移除
        emit CertificateApproved(_id, _address); // 触发事件
    }
    
    // 审核拒绝证书
    function rejectCertificate(address _address, uint _id) external {
        require(keccak256(abi.encodePacked(certificates[_id].status)) == keccak256(abi.encodePacked("Pending")), "Certificate is not pending"); // 断言证书状态为待审核
        require(certificates[_id].approver == address(0), "Certificate already approved or rejected"); // 断言证书未被审核过
        certificates[_id].status = "Rejected"; // 更新证书状态为审核拒绝
        certificates[_id].approvalTime = block.timestamp; // 更新审核时间
        certificates[_id].approver = _address; // 更新审核人
        removePendingCertificate(_id); // 从待审核列表中移除
        emit CertificateRejected(_id, _address); // 触发事件
    }
    
    // 从待审核列表中移除证书
    function removePendingCertificate(uint _id) internal {
        for (uint i = 0; i < pendingCertificates.length; i++) {
            if (pendingCertificates[i] == _id) {
                pendingCertificates[i] = pendingCertificates[pendingCertificates.length - 1];
                pendingCertificates.pop();
                break;
            }
        }
    }
    
    // 根据ID获取证书详情
    function getCertificateDetails(uint _id) external view returns (uint, string memory, uint, string memory, address, uint, uint, string memory, address) {
        Certificate memory cert = certificates[_id];
        return (cert.id, cert.studentName, cert.grade, cert.course, cert.studentAddress, cert.applicationTime, cert.approvalTime, cert.status, cert.approver);
    }
    
    // 获取学生的证书列表, 可选条件：课程 状态
    
    function getStudentCertificates(address _student, string memory _course, string memory _status) external view returns (Certificate[] memory) {
        Certificate[] memory certs = new Certificate[](studentCertificates[_student].length);
        uint count = 0;
        for (uint i = 0; i < studentCertificates[_student].length; i++) {
            Certificate memory cert = certificates[studentCertificates[_student][i]];
            // 判断是否符合条件
            if(bytes(_course).length != 0 && keccak256(abi.encodePacked(_course)) != keccak256(abi.encodePacked(cert.course))){
                continue;
            }
            if(bytes(_status).length != 0 && keccak256(abi.encodePacked(_status)) != keccak256(abi.encodePacked(cert.status))){
                continue;
            }
            certs[count] = cert;
            count++;
        }
        return certs;
    }
    
    // 获取待审核的证书列表, 可选条件：课程 学生姓名 学生地址
    function getPendingCertificates(address _student, string memory _course, string memory _studentName) external view returns (Certificate[] memory) {
        Certificate[] memory certs = new Certificate[](pendingCertificates.length);
        uint count = 0;
        for (uint i = 0; i < pendingCertificates.length; i++) {
            Certificate memory cert = certificates[pendingCertificates[i]];
            // 判断是否符合条件
            if(bytes(_course).length != 0 && keccak256(abi.encodePacked(_course)) != keccak256(abi.encodePacked(cert.course))){
                continue;
            }

            if(bytes(_studentName).length != 0 && keccak256(abi.encodePacked(_studentName)) != keccak256(abi.encodePacked(cert.studentName))){
                continue;
            }
            if(_student != address(0) && _student != cert.studentAddress){
                continue;
            }
            certs[count] = cert;
            count++;
        }
        return certs;
    }
}
