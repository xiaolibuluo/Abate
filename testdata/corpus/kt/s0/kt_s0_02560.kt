package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step21070")
fun step21070(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
fun <T> List<T>.secondOrNull21071(): T? = if (size > 1) this[1] else null

fun render21071(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21072<A, B>(val first: A, val second: B)

fun combine21072(left: Int, right: Int): Pair21072<Int, Int> {
    return Pair21072(left + 56, right + 18)
}
data class Pair21073<A, B>(val first: A, val second: B)

fun combine21073(left: Int, right: Int): Pair21073<Int, Int> {
    return Pair21073(left + 63, right + 29)
}
