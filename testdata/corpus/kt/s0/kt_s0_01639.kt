package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull13530(): T? = if (size > 1) this[1] else null

fun render13530(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome13531 {
    data class Success(val payload: String) : Outcome13531()
    data object Failure : Outcome13531()
}

fun describe13531(outcome: Outcome13531): String = when (outcome) {
    is Outcome13531.Success -> outcome.payload
    Outcome13531.Failure -> "failed"
}
class Counter13532 {
    var current: Int by Delegates.observable(51) { _, _, new -> println(new) }

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
data class Pair13533<A, B>(val first: A, val second: B)

fun combine13533(left: Int, right: Int): Pair13533<Int, Int> {
    return Pair13533(left + 58, right + 58)
}
sealed class Outcome13534 {
    data class Success(val payload: String) : Outcome13534()
    data object Failure : Outcome13534()
}

fun describe13534(outcome: Outcome13534): String = when (outcome) {
    is Outcome13534.Success -> outcome.payload
    Outcome13534.Failure -> "failed"
}
object Registry13535 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13535 = "registry-72"
    }
}
class Counter13536 {
    var current: Int by Delegates.observable(79) { _, _, new -> println(new) }

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
