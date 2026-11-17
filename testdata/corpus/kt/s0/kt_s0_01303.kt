package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull10750(): T? = if (size > 1) this[1] else null

fun render10750(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair10751<A, B>(val first: A, val second: B)

fun combine10751(left: Int, right: Int): Pair10751<Int, Int> {
    return Pair10751(left + 19, right + 42)
}
