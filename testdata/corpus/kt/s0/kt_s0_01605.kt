package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry13250 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13250 = "registry-61"
    }
}
sealed class Outcome13251 {
    data class Success(val payload: String) : Outcome13251()
    data object Failure : Outcome13251()
}

fun describe13251(outcome: Outcome13251): String = when (outcome) {
    is Outcome13251.Success -> outcome.payload
    Outcome13251.Failure -> "failed"
}
sealed class Outcome13252 {
    data class Success(val payload: String) : Outcome13252()
    data object Failure : Outcome13252()
}

fun describe13252(outcome: Outcome13252): String = when (outcome) {
    is Outcome13252.Success -> outcome.payload
    Outcome13252.Failure -> "failed"
}
