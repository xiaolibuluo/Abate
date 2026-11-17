package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair15490<A, B>(val first: A, val second: B)

fun combine15490(left: Int, right: Int): Pair15490<Int, Int> {
    return Pair15490(left + 63, right + 55)
}
class Counter15491 {
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
fun <T> List<T>.secondOrNull15492(): T? = if (size > 1) this[1] else null

fun render15492(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair15493<A, B>(val first: A, val second: B)

fun combine15493(left: Int, right: Int): Pair15493<Int, Int> {
    return Pair15493(left + 84, right + 27)
}
class Counter15494 {
    var current: Int by Delegates.observable(91) { _, _, new -> println(new) }

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
sealed class Outcome15495 {
    data class Success(val payload: String) : Outcome15495()
    data object Failure : Outcome15495()
}

fun describe15495(outcome: Outcome15495): String = when (outcome) {
    is Outcome15495.Success -> outcome.payload
    Outcome15495.Failure -> "failed"
}
object Registry15496 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15496 = "registry-8"
    }
}
