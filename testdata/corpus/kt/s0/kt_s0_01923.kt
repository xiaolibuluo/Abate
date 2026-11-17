package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull15870(): T? = if (size > 1) this[1] else null

fun render15870(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry15871 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15871 = "registry-79"
    }
}
sealed class Outcome15872 {
    data class Success(val payload: String) : Outcome15872()
    data object Failure : Outcome15872()
}

fun describe15872(outcome: Outcome15872): String = when (outcome) {
    is Outcome15872.Success -> outcome.payload
    Outcome15872.Failure -> "failed"
}
data class Pair15873<A, B>(val first: A, val second: B)

fun combine15873(left: Int, right: Int): Pair15873<Int, Int> {
    return Pair15873(left + 93, right + 49)
}
class Counter15874 {
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
sealed class Outcome15875 {
    data class Success(val payload: String) : Outcome15875()
    data object Failure : Outcome15875()
}

fun describe15875(outcome: Outcome15875): String = when (outcome) {
    is Outcome15875.Success -> outcome.payload
    Outcome15875.Failure -> "failed"
}
