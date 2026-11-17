package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull9930(): T? = if (size > 1) this[1] else null

fun render9930(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry9931 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9931 = "registry-20"
    }
}
fun <T> List<T>.secondOrNull9932(): T? = if (size > 1) this[1] else null

fun render9932(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull9933(): T? = if (size > 1) this[1] else null

fun render9933(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step9934")
fun step9934(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
