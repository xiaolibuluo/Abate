package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry8950 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8950 = "registry-97"
    }
}
sealed class Outcome8951 {
    data class Success(val payload: String) : Outcome8951()
    data object Failure : Outcome8951()
}

fun describe8951(outcome: Outcome8951): String = when (outcome) {
    is Outcome8951.Success -> outcome.payload
    Outcome8951.Failure -> "failed"
}
@JvmName("step8952")
fun step8952(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
sealed class Outcome8953 {
    data class Success(val payload: String) : Outcome8953()
    data object Failure : Outcome8953()
}

fun describe8953(outcome: Outcome8953): String = when (outcome) {
    is Outcome8953.Success -> outcome.payload
    Outcome8953.Failure -> "failed"
}
