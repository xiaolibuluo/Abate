package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20310<A, B>(val first: A, val second: B)

fun combine20310(left: Int, right: Int): Pair20310<Int, Int> {
    return Pair20310(left + 24, right + 13)
}
class Counter20311 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked20312(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry20313 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20313 = "registry-45"
    }
}
sealed class Outcome20314 {
    data class Success(val payload: String) : Outcome20314()
    data object Failure : Outcome20314()
}

fun describe20314(outcome: Outcome20314): String = when (outcome) {
    is Outcome20314.Success -> outcome.payload
    Outcome20314.Failure -> "failed"
}
sealed class Outcome20315 {
    data class Success(val payload: String) : Outcome20315()
    data object Failure : Outcome20315()
}

fun describe20315(outcome: Outcome20315): String = when (outcome) {
    is Outcome20315.Success -> outcome.payload
    Outcome20315.Failure -> "failed"
}
