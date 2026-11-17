package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull4310(): T? = if (size > 1) this[1] else null

fun render4310(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair4311<A, B>(val first: A, val second: B)

fun combine4311(left: Int, right: Int): Pair4311<Int, Int> {
    return Pair4311(left + 86, right + 48)
}
class Counter4312 {
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
fun <T> List<T>.secondOrNull4313(): T? = if (size > 1) this[1] else null

fun render4313(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull4314(): T? = if (size > 1) this[1] else null

fun render4314(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome4315 {
    data class Success(val payload: String) : Outcome4315()
    data object Failure : Outcome4315()
}

fun describe4315(outcome: Outcome4315): String = when (outcome) {
    is Outcome4315.Success -> outcome.payload
    Outcome4315.Failure -> "failed"
}
object Registry4316 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4316 = "registry-24"
    }
}
data class Pair4317<A, B>(val first: A, val second: B)

fun combine4317(left: Int, right: Int): Pair4317<Int, Int> {
    return Pair4317(left + 31, right + 53)
}
