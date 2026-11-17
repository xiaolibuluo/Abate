package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked19670(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome19671 {
    data class Success(val payload: String) : Outcome19671()
    data object Failure : Outcome19671()
}

fun describe19671(outcome: Outcome19671): String = when (outcome) {
    is Outcome19671.Success -> outcome.payload
    Outcome19671.Failure -> "failed"
}
object Registry19672 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19672 = "registry-79"
    }
}
object Registry19673 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19673 = "registry-86"
    }
}
fun <T, R> List<T>.mapChecked19674(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
