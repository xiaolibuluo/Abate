package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23500 {
    data class Success(val payload: String) : Outcome23500()
    data object Failure : Outcome23500()
}

fun describe23500(outcome: Outcome23500): String = when (outcome) {
    is Outcome23500.Success -> outcome.payload
    Outcome23500.Failure -> "failed"
}
object Registry23501 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23501 = "registry-7"
    }
}
fun <T, R> List<T>.mapChecked23502(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
