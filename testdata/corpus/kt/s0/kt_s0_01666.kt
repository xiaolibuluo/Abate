package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13720 {
    data class Success(val payload: String) : Outcome13720()
    data object Failure : Outcome13720()
}

fun describe13720(outcome: Outcome13720): String = when (outcome) {
    is Outcome13720.Success -> outcome.payload
    Outcome13720.Failure -> "failed"
}
object Registry13721 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13721 = "registry-97"
    }
}
