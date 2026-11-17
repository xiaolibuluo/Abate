package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2370 {
    data class Success(val payload: String) : Outcome2370()
    data object Failure : Outcome2370()
}

fun describe2370(outcome: Outcome2370): String = when (outcome) {
    is Outcome2370.Success -> outcome.payload
    Outcome2370.Failure -> "failed"
}
class Counter2371 {
    var current: Int by Delegates.observable(86) { _, _, new -> println(new) }

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
data class Pair2372<A, B>(val first: A, val second: B)

fun combine2372(left: Int, right: Int): Pair2372<Int, Int> {
    return Pair2372(left + 93, right + 43)
}
class Counter2373 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
object Registry2374 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2374 = "registry-10"
    }
}
fun <T> List<T>.secondOrNull2375(): T? = if (size > 1) this[1] else null

fun render2375(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome2376 {
    data class Success(val payload: String) : Outcome2376()
    data object Failure : Outcome2376()
}

fun describe2376(outcome: Outcome2376): String = when (outcome) {
    is Outcome2376.Success -> outcome.payload
    Outcome2376.Failure -> "failed"
}
