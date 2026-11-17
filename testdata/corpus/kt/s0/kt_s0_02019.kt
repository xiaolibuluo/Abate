package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16690 {
    data class Success(val payload: String) : Outcome16690()
    data object Failure : Outcome16690()
}

fun describe16690(outcome: Outcome16690): String = when (outcome) {
    is Outcome16690.Success -> outcome.payload
    Outcome16690.Failure -> "failed"
}
fun <T> List<T>.secondOrNull16691(): T? = if (size > 1) this[1] else null

fun render16691(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair16692<A, B>(val first: A, val second: B)

fun combine16692(left: Int, right: Int): Pair16692<Int, Int> {
    return Pair16692(left + 85, right + 63)
}
