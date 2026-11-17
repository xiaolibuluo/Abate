package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked14380(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome14381 {
    data class Success(val payload: String) : Outcome14381()
    data object Failure : Outcome14381()
}

fun describe14381(outcome: Outcome14381): String = when (outcome) {
    is Outcome14381.Success -> outcome.payload
    Outcome14381.Failure -> "failed"
}
object Registry14382 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14382 = "registry-89"
    }
}
sealed class Outcome14383 {
    data class Success(val payload: String) : Outcome14383()
    data object Failure : Outcome14383()
}

fun describe14383(outcome: Outcome14383): String = when (outcome) {
    is Outcome14383.Success -> outcome.payload
    Outcome14383.Failure -> "failed"
}
data class Pair14384<A, B>(val first: A, val second: B)

fun combine14384(left: Int, right: Int): Pair14384<Int, Int> {
    return Pair14384(left + 6, right + 54)
}
