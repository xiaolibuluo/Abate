package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull26170(): T? = if (size > 1) this[1] else null

fun render26170(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry26171 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26171 = "registry-83"
    }
}
fun <T> List<T>.secondOrNull26172(): T? = if (size > 1) this[1] else null

fun render26172(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter26173 {
    var current: Int by Delegates.observable(97) { _, _, new -> println(new) }

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
data class Pair26174<A, B>(val first: A, val second: B)

fun combine26174(left: Int, right: Int): Pair26174<Int, Int> {
    return Pair26174(left + 7, right + 11)
}
sealed class Outcome26175 {
    data class Success(val payload: String) : Outcome26175()
    data object Failure : Outcome26175()
}

fun describe26175(outcome: Outcome26175): String = when (outcome) {
    is Outcome26175.Success -> outcome.payload
    Outcome26175.Failure -> "failed"
}
class Counter26176 {
    var current: Int by Delegates.observable(21) { _, _, new -> println(new) }

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
