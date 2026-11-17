package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry17250 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17250 = "registry-23"
    }
}
sealed class Outcome17251 {
    data class Success(val payload: String) : Outcome17251()
    data object Failure : Outcome17251()
}

fun describe17251(outcome: Outcome17251): String = when (outcome) {
    is Outcome17251.Success -> outcome.payload
    Outcome17251.Failure -> "failed"
}
class Counter17252 {
    var current: Int by Delegates.observable(37) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T, R> List<T>.mapChecked17253(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
