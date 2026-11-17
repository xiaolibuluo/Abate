package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull7110(): T? = if (size > 1) this[1] else null

fun render7110(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome7111 {
    data class Success(val payload: String) : Outcome7111()
    data object Failure : Outcome7111()
}

fun describe7111(outcome: Outcome7111): String = when (outcome) {
    is Outcome7111.Success -> outcome.payload
    Outcome7111.Failure -> "failed"
}
data class Pair7112<A, B>(val first: A, val second: B)

fun combine7112(left: Int, right: Int): Pair7112<Int, Int> {
    return Pair7112(left + 47, right + 67)
}
class Counter7113 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

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
object Registry7114 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7114 = "registry-61"
    }
}
object Registry7115 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7115 = "registry-68"
    }
}
