package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull3070(): T? = if (size > 1) this[1] else null

fun render3070(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair3071<A, B>(val first: A, val second: B)

fun combine3071(left: Int, right: Int): Pair3071<Int, Int> {
    return Pair3071(left + 26, right + 34)
}
fun <T, R> List<T>.mapChecked3072(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair3073<A, B>(val first: A, val second: B)

fun combine3073(left: Int, right: Int): Pair3073<Int, Int> {
    return Pair3073(left + 40, right + 56)
}
