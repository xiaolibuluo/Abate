package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry13120 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13120 = "registry-86"
    }
}
class Counter13121 {
    var current: Int by Delegates.observable(93) { _, _, new -> println(new) }

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
sealed class Outcome13122 {
    data class Success(val payload: String) : Outcome13122()
    data object Failure : Outcome13122()
}

fun describe13122(outcome: Outcome13122): String = when (outcome) {
    is Outcome13122.Success -> outcome.payload
    Outcome13122.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked13123(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome13124 {
    data class Success(val payload: String) : Outcome13124()
    data object Failure : Outcome13124()
}

fun describe13124(outcome: Outcome13124): String = when (outcome) {
    is Outcome13124.Success -> outcome.payload
    Outcome13124.Failure -> "failed"
}
data class Pair13125<A, B>(val first: A, val second: B)

fun combine13125(left: Int, right: Int): Pair13125<Int, Int> {
    return Pair13125(left + 24, right + 37)
}
fun <T> List<T>.secondOrNull13126(): T? = if (size > 1) this[1] else null

fun render13126(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter13127 {
    var current: Int by Delegates.observable(38) { _, _, new -> println(new) }

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
