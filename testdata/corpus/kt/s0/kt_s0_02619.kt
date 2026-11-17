package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step21610")
fun step21610(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
fun <T> List<T>.secondOrNull21611(): T? = if (size > 1) this[1] else null

fun render21611(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21612<A, B>(val first: A, val second: B)

fun combine21612(left: Int, right: Int): Pair21612<Int, Int> {
    return Pair21612(left + 79, right + 30)
}
@JvmName("step21613")
fun step21613(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 86
}
data class Pair21614<A, B>(val first: A, val second: B)

fun combine21614(left: Int, right: Int): Pair21614<Int, Int> {
    return Pair21614(left + 93, right + 52)
}
@JvmName("step21615")
fun step21615(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
