package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16650<A, B>(val first: A, val second: B)

fun combine16650(left: Int, right: Int): Pair16650<Int, Int> {
    return Pair16650(left + 19, right + 13)
}
fun <T> List<T>.secondOrNull16651(): T? = if (size > 1) this[1] else null

fun render16651(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step16652")
fun step16652(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
