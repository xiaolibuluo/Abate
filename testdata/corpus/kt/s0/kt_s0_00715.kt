package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5890(): T? = if (size > 1) this[1] else null

fun render5890(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair5891<A, B>(val first: A, val second: B)

fun combine5891(left: Int, right: Int): Pair5891<Int, Int> {
    return Pair5891(left + 6, right + 56)
}
