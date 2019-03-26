const getCellValue = (tr, idx) => tr.children[idx].innerText || tr.children[idx].textContent;

const comparer = function (idx, asc) {
	return (a, b) => ((v1, v2) => {
		if (v1 == '' || v2 == '' || isNaN(v1) || isNaN(v2)) {
			return v1.toString().localeCompare(v2);
		}
		return v1 - v2;
	})(getCellValue(asc ? a : b, idx), getCellValue(asc ? b : a, idx));
};

document.querySelectorAll('th').forEach(th => th.addEventListener('click', (() => {
	tbls = document.getElementsByTagName('table');
    const thi = Array.from(th.parentNode.children).indexOf(th);

    this.asc = !this.asc;

    if (this.asc) {
        document.getElementsByTagName('html')[0].className = 'sorted-by-' + (thi + 1) + '-asc';
    } else {
        document.getElementsByTagName('html')[0].className = 'sorted-by-' + (thi + 1) + '-desc';
    }

    for (let i = 0; i < tbls.length; i++) {
	    const tbody = tbls[i].querySelector('tbody');
	    Array.from(tbody.querySelectorAll('tr')).sort(comparer(thi, this.asc)).forEach(tr => tbody.appendChild(tr));
    }
})));

