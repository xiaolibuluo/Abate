package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair17070<A, B>(val first: A, val second: B)

fun combine17070(left: Int, right: Int): Pair17070<Int, Int> {
    return Pair17070(left + 80, right + 63)
}
fun <T> List<T>.secondOrNull17071(): T? = if (size > 1) this[1] else null

fun render17071(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
