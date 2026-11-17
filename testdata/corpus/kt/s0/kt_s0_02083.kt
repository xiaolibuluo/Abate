package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull17220(): T? = if (size > 1) this[1] else null

fun render17220(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step17221")
fun step17221(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
@JvmName("step17222")
fun step17222(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
data class Pair17223<A, B>(val first: A, val second: B)

fun combine17223(left: Int, right: Int): Pair17223<Int, Int> {
    return Pair17223(left + 5, right + 18)
}
fun <T> List<T>.secondOrNull17224(): T? = if (size > 1) this[1] else null

fun render17224(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
