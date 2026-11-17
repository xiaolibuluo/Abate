package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked24000(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome24001 {
    data class Success(val payload: String) : Outcome24001()
    data object Failure : Outcome24001()
}

fun describe24001(outcome: Outcome24001): String = when (outcome) {
    is Outcome24001.Success -> outcome.payload
    Outcome24001.Failure -> "failed"
}
object Registry24002 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24002 = "registry-82"
    }
}
object Registry24003 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24003 = "registry-89"
    }
}
