package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked18370(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked18371(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked18372(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter18373 {
    var current: Int by Delegates.observable(45) { _, _, new -> println(new) }

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
object Registry18374 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18374 = "registry-52"
    }
}
sealed class Outcome18375 {
    data class Success(val payload: String) : Outcome18375()
    data object Failure : Outcome18375()
}

fun describe18375(outcome: Outcome18375): String = when (outcome) {
    is Outcome18375.Success -> outcome.payload
    Outcome18375.Failure -> "failed"
}
sealed class Outcome18376 {
    data class Success(val payload: String) : Outcome18376()
    data object Failure : Outcome18376()
}

fun describe18376(outcome: Outcome18376): String = when (outcome) {
    is Outcome18376.Success -> outcome.payload
    Outcome18376.Failure -> "failed"
}
