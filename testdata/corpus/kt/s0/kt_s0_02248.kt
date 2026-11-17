package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter18480 {
    var current: Int by Delegates.observable(70) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked18481(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull18482(): T? = if (size > 1) this[1] else null

fun render18482(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked18483(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter18484 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
data class Pair18485<A, B>(val first: A, val second: B)

fun combine18485(left: Int, right: Int): Pair18485<Int, Int> {
    return Pair18485(left + 8, right + 68)
}
object Registry18486 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18486 = "registry-15"
    }
}
sealed class Outcome18487 {
    data class Success(val payload: String) : Outcome18487()
    data object Failure : Outcome18487()
}

fun describe18487(outcome: Outcome18487): String = when (outcome) {
    is Outcome18487.Success -> outcome.payload
    Outcome18487.Failure -> "failed"
}
