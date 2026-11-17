package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome8480 {
    data class Success(val payload: String) : Outcome8480()
    data object Failure : Outcome8480()
}

fun describe8480(outcome: Outcome8480): String = when (outcome) {
    is Outcome8480.Success -> outcome.payload
    Outcome8480.Failure -> "failed"
}
object Registry8481 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8481 = "registry-75"
    }
}
@JvmName("step8482")
fun step8482(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
sealed class Outcome8483 {
    data class Success(val payload: String) : Outcome8483()
    data object Failure : Outcome8483()
}

fun describe8483(outcome: Outcome8483): String = when (outcome) {
    is Outcome8483.Success -> outcome.payload
    Outcome8483.Failure -> "failed"
}
