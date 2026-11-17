package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull17150(): T? = if (size > 1) this[1] else null

fun render17150(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair17151<A, B>(val first: A, val second: B)

fun combine17151(left: Int, right: Int): Pair17151<Int, Int> {
    return Pair17151(left + 94, right + 69)
}
data class Pair17152<A, B>(val first: A, val second: B)

fun combine17152(left: Int, right: Int): Pair17152<Int, Int> {
    return Pair17152(left + 101, right + 19)
}
data class Pair17153<A, B>(val first: A, val second: B)

fun combine17153(left: Int, right: Int): Pair17153<Int, Int> {
    return Pair17153(left + 11, right + 30)
}
