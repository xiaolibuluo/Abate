package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6540 {
    data class Success(val payload: String) : Outcome6540()
    data object Failure : Outcome6540()
}

fun describe6540(outcome: Outcome6540): String = when (outcome) {
    is Outcome6540.Success -> outcome.payload
    Outcome6540.Failure -> "failed"
}
sealed class Outcome6541 {
    data class Success(val payload: String) : Outcome6541()
    data object Failure : Outcome6541()
}

fun describe6541(outcome: Outcome6541): String = when (outcome) {
    is Outcome6541.Success -> outcome.payload
    Outcome6541.Failure -> "failed"
}
object Registry6542 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6542 = "registry-82"
    }
}
@JvmName("step6543")
fun step6543(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
