package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6960")
fun step6960(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
sealed class Outcome6961 {
    data class Success(val payload: String) : Outcome6961()
    data object Failure : Outcome6961()
}

fun describe6961(outcome: Outcome6961): String = when (outcome) {
    is Outcome6961.Success -> outcome.payload
    Outcome6961.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked6962(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry6963 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6963 = "registry-53"
    }
}
