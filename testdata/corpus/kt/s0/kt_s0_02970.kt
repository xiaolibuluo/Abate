package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step24430")
fun step24430(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
data class Pair24431<A, B>(val first: A, val second: B)

fun combine24431(left: Int, right: Int): Pair24431<Int, Int> {
    return Pair24431(left + 52, right + 41)
}
@JvmName("step24432")
fun step24432(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 59
}
fun <T> List<T>.secondOrNull24433(): T? = if (size > 1) this[1] else null

fun render24433(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
