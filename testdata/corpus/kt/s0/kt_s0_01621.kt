package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry13380 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13380 = "registry-36"
    }
}
@JvmName("step13381")
fun step13381(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
sealed class Outcome13382 {
    data class Success(val payload: String) : Outcome13382()
    data object Failure : Outcome13382()
}

fun describe13382(outcome: Outcome13382): String = when (outcome) {
    is Outcome13382.Success -> outcome.payload
    Outcome13382.Failure -> "failed"
}
sealed class Outcome13383 {
    data class Success(val payload: String) : Outcome13383()
    data object Failure : Outcome13383()
}

fun describe13383(outcome: Outcome13383): String = when (outcome) {
    is Outcome13383.Success -> outcome.payload
    Outcome13383.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13384(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
