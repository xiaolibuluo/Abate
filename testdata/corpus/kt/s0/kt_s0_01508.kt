package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter12480 {
    var current: Int by Delegates.observable(30) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked12481(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12482 {
    data class Success(val payload: String) : Outcome12482()
    data object Failure : Outcome12482()
}

fun describe12482(outcome: Outcome12482): String = when (outcome) {
    is Outcome12482.Success -> outcome.payload
    Outcome12482.Failure -> "failed"
}
object Registry12483 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12483 = "registry-51"
    }
}
fun <T, R> List<T>.mapChecked12484(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry12485 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12485 = "registry-65"
    }
}
