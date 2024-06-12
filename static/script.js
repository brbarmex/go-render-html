document.addEventListener("DOMContentLoaded", function() {
    fetch('/api/auditorias')
        .then(response => response.json())
        .then(data => {
            const tableBody = document.querySelector("#auditoriaTable tbody");
            data.forEach(auditoria => {
                const row = document.createElement("tr");

                const versaoCell = document.createElement("td");
                versaoCell.textContent = auditoria.versao;
                row.appendChild(versaoCell);

                const dataHoraCell = document.createElement("td");
                dataHoraCell.textContent = new Date(auditoria.data_hora).toLocaleString();
                row.appendChild(dataHoraCell);

                const executorCell = document.createElement("td");
                executorCell.textContent = auditoria.executor;
                row.appendChild(executorCell);

                tableBody.appendChild(row);
            });
        })
        .catch(error => console.error('Erro ao buscar auditorias:', error));
});
