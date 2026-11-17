package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter22790 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
sealed class Outcome22791 {
    data class Success(val payload: String) : Outcome22791()
    data object Failure : Outcome22791()
}

fun describe22791(outcome: Outcome22791): String = when (outcome) {
    is Outcome22791.Success -> outcome.payload
    Outcome22791.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked22792(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry22793 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22793 = "registry-68"
    }
}
fun <T, R> List<T>.mapChecked22794(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry22795 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22795 = "registry-82"
    }
}
object Registry22796 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22796 = "registry-89"
    }
}
sealed class Outcome22797 {
    data class Success(val payload: String) : Outcome22797()
    data object Failure : Outcome22797()
}

fun describe22797(outcome: Outcome22797): String = when (outcome) {
    is Outcome22797.Success -> outcome.payload
    Outcome22797.Failure -> "failed"
}
