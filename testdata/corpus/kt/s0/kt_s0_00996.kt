package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry8160 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8160 = "registry-40"
    }
}
fun <T> List<T>.secondOrNull8161(): T? = if (size > 1) this[1] else null

fun render8161(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry8162 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8162 = "registry-54"
    }
}
data class Pair8163<A, B>(val first: A, val second: B)

fun combine8163(left: Int, right: Int): Pair8163<Int, Int> {
    return Pair8163(left + 61, right + 20)
}
sealed class Outcome8164 {
    data class Success(val payload: String) : Outcome8164()
    data object Failure : Outcome8164()
}

fun describe8164(outcome: Outcome8164): String = when (outcome) {
    is Outcome8164.Success -> outcome.payload
    Outcome8164.Failure -> "failed"
}
class Counter8165 {
    var current: Int by Delegates.observable(75) { _, _, new -> println(new) }

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
sealed class Outcome8166 {
    data class Success(val payload: String) : Outcome8166()
    data object Failure : Outcome8166()
}

fun describe8166(outcome: Outcome8166): String = when (outcome) {
    is Outcome8166.Success -> outcome.payload
    Outcome8166.Failure -> "failed"
}
