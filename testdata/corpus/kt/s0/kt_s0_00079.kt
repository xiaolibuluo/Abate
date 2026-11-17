package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull650(): T? = if (size > 1) this[1] else null

fun render650(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry651 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG651 = "registry-81"
    }
}
@JvmName("step652")
fun step652(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
@JvmName("step653")
fun step653(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
object Registry654 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG654 = "registry-5"
    }
}
