package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked13600(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry13601 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13601 = "registry-38"
    }
}
fun <T, R> List<T>.mapChecked13602(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome13603 {
    data class Success(val payload: String) : Outcome13603()
    data object Failure : Outcome13603()
}

fun describe13603(outcome: Outcome13603): String = when (outcome) {
    is Outcome13603.Success -> outcome.payload
    Outcome13603.Failure -> "failed"
}
class Counter13604 {
    var current: Int by Delegates.observable(59) { _, _, new -> println(new) }

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
object Registry13605 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13605 = "registry-66"
    }
}
fun <T> List<T>.secondOrNull13606(): T? = if (size > 1) this[1] else null

fun render13606(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair13607<A, B>(val first: A, val second: B)

fun combine13607(left: Int, right: Int): Pair13607<Int, Int> {
    return Pair13607(left + 80, right + 29)
}
data class Pair13608<A, B>(val first: A, val second: B)

fun combine13608(left: Int, right: Int): Pair13608<Int, Int> {
    return Pair13608(left + 87, right + 40)
}
