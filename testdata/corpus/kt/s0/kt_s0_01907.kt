package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15730 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15730 = "registry-84"
    }
}
object Registry15731 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15731 = "registry-91"
    }
}
sealed class Outcome15732 {
    data class Success(val payload: String) : Outcome15732()
    data object Failure : Outcome15732()
}

fun describe15732(outcome: Outcome15732): String = when (outcome) {
    is Outcome15732.Success -> outcome.payload
    Outcome15732.Failure -> "failed"
}
@JvmName("step15733")
fun step15733(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
