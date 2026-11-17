package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull13330(): T? = if (size > 1) this[1] else null

fun render13330(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry13331 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13331 = "registry-75"
    }
}
data class Pair13332<A, B>(val first: A, val second: B)

fun combine13332(left: Int, right: Int): Pair13332<Int, Int> {
    return Pair13332(left + 82, right + 29)
}
fun <T> List<T>.secondOrNull13333(): T? = if (size > 1) this[1] else null

fun render13333(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step13334")
fun step13334(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
