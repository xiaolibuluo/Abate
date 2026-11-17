package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12720 {
    data class Success(val payload: String) : Outcome12720()
    data object Failure : Outcome12720()
}

fun describe12720(outcome: Outcome12720): String = when (outcome) {
    is Outcome12720.Success -> outcome.payload
    Outcome12720.Failure -> "failed"
}
object Registry12721 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12721 = "registry-58"
    }
}
