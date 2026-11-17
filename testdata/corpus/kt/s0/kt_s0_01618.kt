package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13350<A, B>(val first: A, val second: B)

fun combine13350(left: Int, right: Int): Pair13350<Int, Int> {
    return Pair13350(left + 94, right + 21)
}
object Registry13351 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13351 = "registry-101"
    }
}
data class Pair13352<A, B>(val first: A, val second: B)

fun combine13352(left: Int, right: Int): Pair13352<Int, Int> {
    return Pair13352(left + 11, right + 43)
}
data class Pair13353<A, B>(val first: A, val second: B)

fun combine13353(left: Int, right: Int): Pair13353<Int, Int> {
    return Pair13353(left + 18, right + 54)
}
class Counter13354 {
    var current: Int by Delegates.observable(25) { _, _, new -> println(new) }

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
