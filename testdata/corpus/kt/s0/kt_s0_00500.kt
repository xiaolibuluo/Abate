package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked4080(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step4081")
fun step4081(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}
object Registry4082 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4082 = "registry-85"
    }
}
sealed class Outcome4083 {
    data class Success(val payload: String) : Outcome4083()
    data object Failure : Outcome4083()
}

fun describe4083(outcome: Outcome4083): String = when (outcome) {
    is Outcome4083.Success -> outcome.payload
    Outcome4083.Failure -> "failed"
}
