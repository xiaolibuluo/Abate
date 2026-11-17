package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9350<A, B>(val first: A, val second: B)

fun combine9350(left: Int, right: Int): Pair9350<Int, Int> {
    return Pair9350(left + 35, right + 27)
}
fun <T> List<T>.secondOrNull9351(): T? = if (size > 1) this[1] else null

fun render9351(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step9352")
fun step9352(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 49
}
