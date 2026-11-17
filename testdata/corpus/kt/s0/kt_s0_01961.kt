package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry16200 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16200 = "registry-16"
    }
}
fun <T, R> List<T>.mapChecked16201(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome16202 {
    data class Success(val payload: String) : Outcome16202()
    data object Failure : Outcome16202()
}

fun describe16202(outcome: Outcome16202): String = when (outcome) {
    is Outcome16202.Success -> outcome.payload
    Outcome16202.Failure -> "failed"
}
