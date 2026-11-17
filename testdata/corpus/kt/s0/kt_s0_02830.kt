package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull23340(): T? = if (size > 1) this[1] else null

fun render23340(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull23341(): T? = if (size > 1) this[1] else null

fun render23341(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter23342 {
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
object Registry23343 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23343 = "registry-7"
    }
}
fun <T> List<T>.secondOrNull23344(): T? = if (size > 1) this[1] else null

fun render23344(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull23345(): T? = if (size > 1) this[1] else null

fun render23345(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair23346<A, B>(val first: A, val second: B)

fun combine23346(left: Int, right: Int): Pair23346<Int, Int> {
    return Pair23346(left + 28, right + 65)
}
sealed class Outcome23347 {
    data class Success(val payload: String) : Outcome23347()
    data object Failure : Outcome23347()
}

fun describe23347(outcome: Outcome23347): String = when (outcome) {
    is Outcome23347.Success -> outcome.payload
    Outcome23347.Failure -> "failed"
}
class Counter23348 {
    var current: Int by Delegates.observable(42) { _, _, new -> println(new) }

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
