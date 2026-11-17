package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry16250 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16250 = "registry-81"
    }
}
sealed class Outcome16251 {
    data class Success(val payload: String) : Outcome16251()
    data object Failure : Outcome16251()
}

fun describe16251(outcome: Outcome16251): String = when (outcome) {
    is Outcome16251.Success -> outcome.payload
    Outcome16251.Failure -> "failed"
}
sealed class Outcome16252 {
    data class Success(val payload: String) : Outcome16252()
    data object Failure : Outcome16252()
}

fun describe16252(outcome: Outcome16252): String = when (outcome) {
    is Outcome16252.Success -> outcome.payload
    Outcome16252.Failure -> "failed"
}
data class Pair16253<A, B>(val first: A, val second: B)

fun combine16253(left: Int, right: Int): Pair16253<Int, Int> {
    return Pair16253(left + 5, right + 10)
}
