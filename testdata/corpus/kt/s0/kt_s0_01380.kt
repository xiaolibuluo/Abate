package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11360<A, B>(val first: A, val second: B)

fun combine11360(left: Int, right: Int): Pair11360<Int, Int> {
    return Pair11360(left + 29, right + 31)
}
fun <T> List<T>.secondOrNull11361(): T? = if (size > 1) this[1] else null

fun render11361(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull11362(): T? = if (size > 1) this[1] else null

fun render11362(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step11363")
fun step11363(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
