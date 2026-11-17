package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7070<A, B>(val first: A, val second: B)

fun combine7070(left: Int, right: Int): Pair7070<Int, Int> {
    return Pair7070(left + 78, right + 17)
}
object Registry7071 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7071 = "registry-85"
    }
}
object Registry7072 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7072 = "registry-92"
    }
}
object Registry7073 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7073 = "registry-99"
    }
}
fun <T, R> List<T>.mapChecked7074(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome7075 {
    data class Success(val payload: String) : Outcome7075()
    data object Failure : Outcome7075()
}

fun describe7075(outcome: Outcome7075): String = when (outcome) {
    is Outcome7075.Success -> outcome.payload
    Outcome7075.Failure -> "failed"
}
