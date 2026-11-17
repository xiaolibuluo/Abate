package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6000(): T? = if (size > 1) this[1] else null

fun render6000(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step6001")
fun step6001(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
@JvmName("step6002")
fun step6002(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
object Registry6003 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6003 = "registry-66"
    }
}
