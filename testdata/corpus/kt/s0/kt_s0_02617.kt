package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21590 {
    data class Success(val payload: String) : Outcome21590()
    data object Failure : Outcome21590()
}

fun describe21590(outcome: Outcome21590): String = when (outcome) {
    is Outcome21590.Success -> outcome.payload
    Outcome21590.Failure -> "failed"
}
data class Pair21591<A, B>(val first: A, val second: B)

fun combine21591(left: Int, right: Int): Pair21591<Int, Int> {
    return Pair21591(left + 46, right + 66)
}
fun <T, R> List<T>.mapChecked21592(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome21593 {
    data class Success(val payload: String) : Outcome21593()
    data object Failure : Outcome21593()
}

fun describe21593(outcome: Outcome21593): String = when (outcome) {
    is Outcome21593.Success -> outcome.payload
    Outcome21593.Failure -> "failed"
}
