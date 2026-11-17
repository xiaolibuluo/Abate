package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5900(): T? = if (size > 1) this[1] else null

fun render5900(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair5901<A, B>(val first: A, val second: B)

fun combine5901(left: Int, right: Int): Pair5901<Int, Int> {
    return Pair5901(left + 19, right + 63)
}
