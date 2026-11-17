package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16190<A, B>(val first: A, val second: B)

fun combine16190(left: Int, right: Int): Pair16190<Int, Int> {
    return Pair16190(left + 100, right + 57)
}
sealed class Outcome16191 {
    data class Success(val payload: String) : Outcome16191()
    data object Failure : Outcome16191()
}

fun describe16191(outcome: Outcome16191): String = when (outcome) {
    is Outcome16191.Success -> outcome.payload
    Outcome16191.Failure -> "failed"
}
fun <T> List<T>.secondOrNull16192(): T? = if (size > 1) this[1] else null

fun render16192(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
