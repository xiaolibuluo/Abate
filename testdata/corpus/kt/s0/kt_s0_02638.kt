package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21790 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
sealed class Outcome21791 {
    data class Success(val payload: String) : Outcome21791()
    data object Failure : Outcome21791()
}

fun describe21791(outcome: Outcome21791): String = when (outcome) {
    is Outcome21791.Success -> outcome.payload
    Outcome21791.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21792(): T? = if (size > 1) this[1] else null

fun render21792(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21793<A, B>(val first: A, val second: B)

fun combine21793(left: Int, right: Int): Pair21793<Int, Int> {
    return Pair21793(left + 29, right + 45)
}
object Registry21794 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21794 = "registry-36"
    }
}
object Registry21795 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21795 = "registry-43"
    }
}
fun <T> List<T>.secondOrNull21796(): T? = if (size > 1) this[1] else null

fun render21796(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
