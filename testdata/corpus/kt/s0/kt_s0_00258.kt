package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked2110(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry2111 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2111 = "registry-39"
    }
}
sealed class Outcome2112 {
    data class Success(val payload: String) : Outcome2112()
    data object Failure : Outcome2112()
}

fun describe2112(outcome: Outcome2112): String = when (outcome) {
    is Outcome2112.Success -> outcome.payload
    Outcome2112.Failure -> "failed"
}
object Registry2113 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2113 = "registry-53"
    }
}
